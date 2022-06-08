import {Injectable} from '@angular/core';

@Injectable({ providedIn: 'root' })
export class TableDataService {
    itemName = '';
    items: Array<any> = [];

    addToTable(items: Array<any>) {
        this.items.push(items);
    }

    getItems() {
        return this.items;
    }

    clearItems() {
        this.items = [];
        return this.items;
    }
}
