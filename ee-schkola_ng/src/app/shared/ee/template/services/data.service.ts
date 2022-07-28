import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public itemIndex: number;
    itemName = '';
    dataName = '';
    items: Map<any, any> = new Map();
    tableItems: Array<any> = [];

    constructor(private _router: Router, private _route: ActivatedRoute) {
    }

    addItemToTableArray(items: Object, id: string) {
        this.items = this.retrieveItemFromCache();
        this.items.set(id, items);
        this.tableItems = this.retrieveItemForTable();
        this.tableItems.push(this.changeObjectToArray(items));
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        localStorage.setItem(this.dataName, JSON.stringify(this.tableItems));
    }

    retrieveItemFromCache() {
        this.items = new Map(JSON.parse(localStorage.map));
        return this.items;
    }

    retrieveItemForTable() {
        this.tableItems = JSON.parse(localStorage.getItem(this.dataName));
        if (this.tableItems !== null) {
            return this.tableItems;
        } else {
            return [];
        }
    }

    changeMapToArray(map: Map<any, any>) {
        console.log(map);
        const tableItems: Array<any> = [];
        map.forEach((value) => {
            tableItems.push(this.changeObjectToArray(value));
        });
        return tableItems
    }

    changeObjectToArray(object: Object) {
        const element: Object = {};
        Object.keys(object).map((elementIndex) => {
            typeof object[elementIndex] === 'object' ?
                Object.keys(object[elementIndex]).map((elementOfObject) => {
                    element[elementOfObject] = object[elementIndex][elementOfObject];
                }) : element[elementIndex] = object[elementIndex];
        });
        return element
    }

    clearItems() {
        this.items.clear();
        localStorage.setItem(this.dataName, JSON.stringify([]));
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        window.location.reload();
    }

    removeItem(id) {
        this.items = this.retrieveItemFromCache();
        this.items.forEach((value, key) => {
            console.log(this.changeObjectToArray(value));
            console.log(id);
            if (JSON.stringify(this.changeObjectToArray(value)) === JSON.stringify(id)) {
                this.items.delete(key);
            }
        });
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        window.location.reload();
    }

    loadEnumElement(enumElement: any) {
        let tempArray = [];
        Object.keys(enumElement).map((element, index) => {
            tempArray.push(enumElement[index]);
        })
        tempArray = tempArray.filter((item) => item);
        return tempArray;
    }

    editItems(index) {
        this._router.navigate(['edit' , index], {relativeTo: this._route});
    }

    checkRoute(element: Object) {
        const currentUrl = this._router.url;
        currentUrl.substring(currentUrl.lastIndexOf('/') + 1).toLowerCase() !== 'new' ? this.isEdit = true : this.isEdit = false;
        this.itemIndex = Number(currentUrl.substring(currentUrl.lastIndexOf('/') + 1).toLowerCase());
        if (this.isEdit) {
            this.loadElement(this.itemIndex, element);
        }
    }

    loadElement(indexValue: number, element: Object) {
        const item = JSON.parse(localStorage.getItem(this.dataName))[indexValue];
        if (item !== null) {
            Object.keys(element).map((elementIndex) => {
                typeof element[elementIndex] === 'object' ?
                    Object.keys(element[elementIndex]).map((elementOfObject) => {
                        element[elementIndex][elementOfObject] = item[elementOfObject];
                    }) : element[elementIndex] = item[elementIndex];
            });
        }
    }
}
