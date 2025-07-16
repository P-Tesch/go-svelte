export function intField(event: InputEvent): void {
    if (event.data === null || !Number.isNaN(parseInt((event as Event & { data: string }).data))) {
        return;
    }
    event.preventDefault();
}

export function floatField(event: InputEvent): void {
    const element: HTMLInputElement = event.target as HTMLInputElement;
    if (event.data === null || (event.data === "," && element.value.match(/,/) === null) || !Number.isNaN(parseInt((event as Event & { data: string }).data))) {
        return;
    }
    event.preventDefault();
}