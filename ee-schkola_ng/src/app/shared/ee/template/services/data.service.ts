import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {SelectionModel} from '@angular/cdk/collections';
import {MatTableDataSource} from '@angular/material/table';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public isSearch: boolean;
    public itemIndex: number;
    filterValue: string;
    itemName = '';
    items: Map<any, any> = new Map();

    selection = new SelectionModel<any>(true, []);
    dataSources: MatTableDataSource<any>;

    constructor(private _router: Router, private _route: ActivatedRoute) {
    }

    addItemToTableArray(items: Object, id: string) {
        Object.keys(items).map((element) => {
            if(typeof items[element] === 'object') {
                for (const elementkey in items[element]) {
                    if(typeof elementkey === 'string') {
                        items[element] = items[element][elementkey];
                        break;
                    }
                }
            }
        });
        this.items = this.retrieveItemsFromCache();
        this.items.set(id, items);
        this.saveItemToCache(this.items);
    }

    retrieveItemsFromCache() {
        this.items = new Map(JSON.parse(localStorage.getItem(this.itemName)));
        return this.items
    }

    changeMapToArray(map: Map<any, any>) {
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
                element[elementIndex] instanceof Date ? element[elementIndex] = new Date(object[elementIndex]) :
                Object.keys(object[elementIndex]).map((elementOfObject) => {
                    element[elementOfObject] = object[elementIndex][elementOfObject];
                }) : element[elementIndex] = object[elementIndex];
        });
        return element;
    }

    clearMultipleItems(selected: any[]) {
        this.items = this.retrieveItemsFromCache();
        this.items.forEach((value, key) => {
            selected.forEach((selectedItem) => {
                if (JSON.stringify(this.changeObjectToArray(value)) === JSON.stringify(selectedItem)) {
                    this.items.delete(key);
                }
            })
        });
        this.saveItemToCache(this.items);
        window.location.reload();
    }

    clearItems() {
        this.items.clear();
        localStorage.setItem(this.itemName, JSON.stringify([]));
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        window.location.reload();
    }

    removeItem(id) {
        this.items = this.retrieveItemsFromCache();
        this.items.forEach((value, key) => {
            if (JSON.stringify(this.changeObjectToArray(value)) === JSON.stringify(id)) {
                this.items.delete(key);
            }
        });
        this.saveItemToCache(this.items);
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

    searchItems(element: Object, relativePath: string) {
        this._router.navigate([ relativePath + '/search'] );
        localStorage.setItem('search', JSON.stringify(element));
    }

    editItems(index: number, element: Object) {
        this._router.navigate([this._router.url + '/edit' , index]);
        localStorage.setItem('edit', JSON.stringify(element));
    }

    checkRoute(element: Object) {
        const currentUrl = this._router.url;
        currentUrl.includes('edit') ? this.isEdit = true : this.isEdit = false;
        this.itemIndex = Number(currentUrl.substring(currentUrl.lastIndexOf('/') + 1).toLowerCase());
        if (this.isEdit) {
            this.loadElement(this.itemIndex, element);
        }
    }

    checkSearchRoute() {
        const currentUrl = this._router.url;
        currentUrl.includes('search') ? this.isSearch = true : this.isSearch = false;
    }

    loadSearchData() {
        const searchItem = JSON.parse(localStorage.getItem('search'));
        this.dataSources = new MatTableDataSource(this.changeMapToArray(this.retrieveItemsFromCache()));
        this.filterValue = searchItem;
        this.dataSources.filter = this.filterValue;
    }

    loadElement(indexValue: number, element: Object) {
        const editItem = JSON.parse(localStorage.getItem('edit'));
        const elementAsObject = JSON.parse(localStorage.getItem(this.itemName))[indexValue][1];
        if (editItem !== null) {
            Object.keys(elementAsObject).map((elementIndex) => {
                typeof elementAsObject[elementIndex] === 'object' ?
                    elementAsObject[elementIndex] instanceof Date ? element[elementIndex] = new Date(editItem[elementIndex]) :
                            element[elementIndex] = elementAsObject[elementIndex] : element[elementIndex] = editItem[elementIndex];
            });
        }
    }

    saveItemToCache(map: Map<any, any>) {
        localStorage.map = JSON.stringify(Array.from(map.entries()));
        localStorage.setItem(this.itemName, localStorage.map);
    }

    applyFilter(event: Event) {
        const filterValue = (event.target as HTMLInputElement).value;
        this.dataSources.filter = filterValue.trim().toLowerCase();
    }

    allRowsSelected() {
        const totalRowSelected = this.selection.selected.length;
        const totalRow = this.dataSources.data.length;
        return totalRowSelected === totalRow;
    }

    masterToggle() {
        this.allRowsSelected() ?
            this.selection.clear() :
            this.dataSources.data.forEach(element => this.selection.select(element));
    }
}
