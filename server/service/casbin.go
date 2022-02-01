package service

import (
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"server/e"
	"server/global"
	"server/model"
	"sync"
)

type CasbinService struct {
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
	policies       []model.CasbinRule
}

var CasbinApp = new(CasbinService)

const CasbinTableName = "casbin_rule"

func (s *CasbinService) Casbin() *casbin.SyncedEnforcer {
	s.once.Do(func() {
		a, err := adapter.NewAdapterByDB(global.DB)
		if err != nil {
			global.LOG.Panic("casbin init adapter error", zap.Error(err))
			return
		}
		s.syncedEnforcer, err = casbin.NewSyncedEnforcer("./resource/casbin_model.conf", a)
		if err != nil {
			global.LOG.Panic("casbin init enforcer error", zap.Error(err))
			return
		}
	})
	_ = s.syncedEnforcer.LoadPolicy()
	return s.syncedEnforcer
}

func (s *CasbinService) Add(policy model.CasbinRule) {
	s.policies = append(s.policies, policy)
}

func (s *CasbinService) AddRange(policies []model.CasbinRule) {
	s.policies = append(s.policies, policies...)
}

func (s *CasbinService) Refresh() e.Err {
	s.ClearAll()
	c := s.Casbin()

	var ruleText [][]string
	for _, rule := range s.policies {
		ruleText = append(ruleText, []string{rule.Role.String(), rule.Path, rule.Method})
	}

	success, err := c.AddPolicies(ruleText)
	if !success || err != nil {
		global.LOG.Warn("casbin add policy error", zap.Error(err))
		return e.CasbinAddPolicyError
	}
	return e.Success
}

func (s *CasbinService) ClearAll() {
	_ = s.Clear(0, model.RoleGuest.String())
	_ = s.Clear(0, model.RoleUser.String())
	_ = s.Clear(0, model.RoleAdmin.String())
}

func (s *CasbinService) Clear(index int, val ...string) e.Err {
	c := s.Casbin()
	success, err := c.RemoveFilteredPolicy(index, val...)
	if !success || err != nil {
		global.LOG.Info("casbin remove policy error",
			zap.Error(err), zap.Int("index", index), zap.Any("val", val))
		return e.CasbinRemovePolicyError
	}
	return e.Success
}
