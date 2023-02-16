import {Injectable} from '@angular/core';
import {Router} from '@angular/router';
import {SelectionModel} from '@angular/cdk/collections';
import {MatTableDataSource} from '@angular/material/table';

@Injectable({ providedIn: 'root' })
export class TableDataService<T> {

    public isEdit: boolean;
    public isSearch: boolean;

    filterValue: string;
    itemName = '';

    items: Map<string, T> = new Map();
    tableItems: Map<string, T> = new Map();
    selection = new SelectionModel<T>(true, []);
    dataSources: MatTableDataSource<T>;

    constructor(protected _router?: Router) {
    }

    inputElement(element: T) {
        const id = this.itemName + JSON.stringify(element);
        this.addItemToTableArray(element, id);
    }

    addItemToTableArray(items: T, id: string) {
        this.items = this.retrieveItemsFromCache()
        this.items.set(id, items);
        this.saveItemToCache(this.items);
    }

    changeMapToArray(data: Map<string, T>) {
        const tableItems: Array<T> = [];
        data.forEach((value) => {
            tableItems.push(value);
        });
        return tableItems
    }

    clearMultipleItems(selected: T[]) {
        this.items = this.retrieveItemsFromCache();
        selected.forEach((selectedItem) => {
            const id = this.itemName + JSON.stringify(selectedItem);
            this.items.delete(id)
        })
        this.saveItemToCache(this.items);
        window.location.reload();
    }

    removeItem(element: T) {
        const id = this.itemName + JSON.stringify(element)
        this.items = this.retrieveItemsFromCache();
        this.items.delete(id)
        this.saveItemToCache(this.items);
        window.location.reload();
    }

    editItems(index: number, element: T) {
        this._router.navigate([this._router.url + '/edit' , index]);
        localStorage.setItem('edit', JSON.stringify(element));
    }

    checkRoute(element: T) {
        const currentUrl = this._router.url;
        currentUrl.includes('edit') ? this.isEdit = true : this.isEdit = false;
        if (this.isEdit) {
            this.loadElement(element);
        }
    }

    checkSearchRoute() {
        const currentUrl = this._router.url;
        currentUrl.includes('search') ? this.isSearch = true : this.isSearch = false;
    }

    loadElement(element: T) {
        const editItem = JSON.parse(localStorage.getItem('edit'));
        if (editItem !== null) {
            Object.assign(element, editItem)
        }
    }

    loadSearchData() {
        const searchItem = JSON.parse(localStorage.getItem('search'));
        this.dataSources = new MatTableDataSource(this.changeMapToArray(this.retrieveItemsFromCache()));
        this.filterValue = searchItem;
        this.dataSources.filter = this.filterValue;
    }

    loadEnumElement(enumElement: any, parentName?: string, elementName?: string) {
        const tempArray = [];
        Object.keys(enumElement).map((element, index) => {
            if (enumElement[index] !== undefined) {
                tempArray.push(parentName + '.' + elementName + '.' + enumElement[index]);
            }
        })
        return tempArray;
    }

    saveItemToCache(data: Map<string, T>) {
        localStorage.map = JSON.stringify(Array.from(data.entries()));
        localStorage.setItem(this.itemName, localStorage.map);
    }

    searchItems(index: number, element: T, relativePath: string, itemName: string) {
        this._router.navigate([ relativePath + '/edit', index] );
        localStorage.setItem('edit', JSON.stringify(element));
        localStorage.setItem('edit-entity', itemName);
    }

    retrieveItemsFromCache(itemName?): Map<string, T> {
        return new Map(JSON.parse(localStorage.getItem(itemName? itemName : this.itemName)));
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

declare global {
    interface Window {
        tableDataService: TableDataService<any>;
    }
}
window.tableDataService = new TableDataService();
