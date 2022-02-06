function processArray<T>(items: T[], process: (item: T) => void) {
    const todo = items.concat();

    setTimeout(function () {
        const item = todo.shift();
        if (item !== undefined) process(item);

        if (todo.length > 0) {
            setTimeout(arguments.callee, 25);
        }
    }, 25);
}

export { processArray };
