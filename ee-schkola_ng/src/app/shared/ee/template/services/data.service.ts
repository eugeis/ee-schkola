import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
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

    selection = new SelectionModel<any>(true, []);
    dataSources: MatTableDataSource<any>;

    constructor(protected _router?: Router) {
    }

    retrieveItemsFromCache() {
        this.items = new Map(JSON.parse(localStorage.getItem(this.itemName)));
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

    /*editInheritedEntity(itemName: string, newElement: any) {
        this.itemName = itemName
        this.items = this.retrieveItemsFromCache();
        const editItem = JSON.parse(localStorage.getItem('edit'));

        if (JSON.stringify(newElement) !== JSON.stringify(editItem)) {
            this.items.forEach((currentMapValue, currentMapKey) => {
                Object.keys(currentMapValue).map((elementIndex) => {
                    if(JSON.stringify(currentMapValue[elementIndex]) === JSON.stringify(editItem)) {
                        newElement = this.changeObjectToArray(this.changeObjectFormat(newElement, this.entityElements));
                        currentMapValue[elementIndex] = newElement
                        const newId = this.itemName + JSON.stringify(currentMapValue);
                        this.items.delete(currentMapKey);
                        this.items.set(newId, currentMapValue);
                        this.saveItemToCache(this.items);
                    } else if (JSON.stringify(currentMapValue[elementIndex]).includes(JSON.stringify(editItem))) {
                        if (typeof currentMapValue[elementIndex] === 'object') {
                            Object.keys(currentMapValue[elementIndex]).map((childIndex) => {
                                if (typeof currentMapValue[elementIndex][childIndex] === 'object') {
                                    currentMapValue[elementIndex][childIndex] = newElement
                                }
                            })
                        }
                        const newId = this.itemName + JSON.stringify(currentMapValue);
                        this.items.delete(currentMapKey);
                        this.items.set(newId, currentMapValue);
                        this.saveItemToCache(this.items);
                    }
                });
            });
        }
    }*/

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
