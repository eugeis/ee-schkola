import {Injectable} from '@angular/core';

@Injectable({ providedIn: 'root' })
export class TableDataService {
    itemName = '';
    items: Array<any> = [];

    addItemToTableArray(items: Array<any>) {
        this.retrieveItemFromCache();
        this.items.push(items);
        localStorage.setItem(this.itemName, JSON.stringify(this.items));
    }

    retrieveItemFromCache() {
        this.items = JSON.parse(localStorage.getItem(this.itemName));
        return this.items;
    }

    getItems() {
        return this.items;
    }

    clearItems() {
        localStorage.setItem(this.itemName, JSON.stringify([]));
        this.items = [];
    }
}
