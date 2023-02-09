import {Injectable} from '@angular/core';
import {Router} from '@angular/router';
import {SelectionModel} from '@angular/cdk/collections';
import {MatTableDataSource} from '@angular/material/table';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public isSearch: boolean;

    filterValue: string;
    entityElements = [];
    itemName = '';
    items: Map<any, any> = new Map();
    tableItems: Map<any, any> = new Map();

    datePattern: string;

    selection = new SelectionModel<any>(true, []);
    dataSources: MatTableDataSource<any>;

    constructor(protected _router?: Router) {
    }

    retrieveItemsFromCache(itemName?) {
        this.items = new Map(JSON.parse(localStorage.getItem(itemName? itemName : this.itemName)));
        return this.items
    }

    clearItems() {
        this.items.clear();
        localStorage.setItem(this.itemName, JSON.stringify([]));
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        window.location.reload();
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

    checkSearchRoute() {
        const currentUrl = this._router.url;
        currentUrl.includes('search') ? this.isSearch = true : this.isSearch = false;
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

    removeSymbolFromString(string: string) {
        return string.replace(/\\/g, '').replace(/"/g, '')
    }

    changeDateValue(element: any) {
        Object.keys(element).map((elementIndex) => {
            if (element[elementIndex] instanceof Date) {
                element[elementIndex] = this.formatDate(element[elementIndex]);
            }
        });
        return element;
    }

    formatDate(date: Date) {
        return [date.getFullYear(), (date.getMonth() + 1), date.getDate()].join('-');
    }
}
