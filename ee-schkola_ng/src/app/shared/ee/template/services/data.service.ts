import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public itemIndex: number;
    itemName = '';
    items: Map<any, any> = new Map();
    tableItems: Array<any> = [];

    constructor(private _router: Router, private _route: ActivatedRoute) {
    }

    addItemToTableArray(items: Object, id: string) {
        this.items = this.retrieveItemFromCache();
        this.items.set(id, items);
        this.tableItems = this.retrieveItemForTable();
        this.tableItems.push(items);
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        localStorage.setItem(this.itemName + 'Data', JSON.stringify(this.tableItems));
    }

    retrieveItemForTable() {
        let element: Object;
        const tableArray = [];
        this.tableItems = JSON.parse(localStorage.getItem(this.itemName + 'Data'));
        if (this.tableItems !== null) {
            Object.keys(this.tableItems).map((itemIndex) => {
                element = {};
                Object.keys(this.tableItems[itemIndex]).map((elementIndex) => {
                    typeof this.tableItems[itemIndex][elementIndex] === 'object' ?
                        Object.keys(this.tableItems[itemIndex][elementIndex]).map((elementOfObject) => {
                            element[elementOfObject] = this.tableItems[itemIndex][elementIndex][elementOfObject];
                        }) : element[elementIndex] = this.tableItems[itemIndex][elementIndex] ;
                });
                tableArray.push(element);
            });
            return tableArray;
        } else {
            return [];
        }
    }

    retrieveItemFromCache() {
        this.items = new Map(JSON.parse(localStorage.map));
        return this.items;
    }

    clearItems() {
        this.items.clear();
        localStorage.setItem(this.itemName + 'Data', JSON.stringify([]));
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        window.location.reload();
    }

    removeItem(id) {
        this.retrieveItemFromCache();
        this.items.delete(id);
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        window.location.reload();
    }

    editItems(index) {
        this._router.navigate(['edit' , index], {relativeTo: this._route});
    }

    loadEnumElement(enumElement: any) {
        let tempArray = [];
        Object.keys(enumElement).map((element, index) => {
            tempArray.push(enumElement[index]);
        })
        tempArray = tempArray.filter((item) => item);
        return tempArray;
    }
}
