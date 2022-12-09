import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {SelectionModel} from '@angular/cdk/collections';
import {MatTableDataSource} from '@angular/material/table';
import {type} from 'os';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public isSearch: boolean;
    public itemIndex: number;
    filterValue: string;
    itemName = '';
    items: Map<any, any> = new Map();
    itemsCopy: Map<any, any> = new Map();

    selection = new SelectionModel<any>(true, []);
    dataSources: MatTableDataSource<any>;

    constructor(private _router: Router, private _route: ActivatedRoute) {
    }

    addItemToTableArray(items: Object, id: string, entityElements: Array<string>) {
        this.items = this.retrieveItemsFromCache();
        this.copyItem(items, id);
        this.changeObjectFormat(items, entityElements);
        this.items.set(id, items);
        this.saveItemToCache(this.items);
    }

    changeObjectFormat(items: Object, entityElements: Array<string>) {
        Object.keys(items).map((itemIndex) => {
            if(typeof items[itemIndex] === 'object') {
                Object.keys(items[itemIndex]).map((childIndex) => {
                    if(entityElements.includes(childIndex)) {
                        items[itemIndex][childIndex] = JSON.stringify(items[itemIndex][childIndex]);
                    }
                })
                if(entityElements.includes(itemIndex)) {
                    items[itemIndex] = JSON.stringify(items[itemIndex]);
                } else {
                    items[itemIndex] = this.changeObjectToArray(items[itemIndex])
                }
            }
        });
    }

    copyItem(items: Object, id: string) {
        this.itemsCopy = this.retrieveCopyItemsFromCache();
        this.itemsCopy.set(id, items);
        this.saveCopyItems(this.itemsCopy);
    }

    retrieveCopyItemsFromCache() {
        const stringCopy = 'items-copy_' + this.itemName;
        this.itemsCopy = new Map(JSON.parse(localStorage.getItem(stringCopy)))
        return this.itemsCopy
    }

    saveCopyItems(map: Map<any, any>) {
        const stringCopy = 'items-copy_' + this.itemName;
        localStorage.copy = JSON.stringify(Array.from(map.entries()));
        localStorage.setItem(stringCopy, localStorage.copy);
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
        Object.keys(object).map((objectIndex) => {
            typeof object[objectIndex] === 'object' ?
                element[objectIndex] instanceof Date ? element[objectIndex] = new Date(object[objectIndex]) :
                Object.keys(object[objectIndex]).map((elementOfObject) => {
                    element[elementOfObject] = object[objectIndex][elementOfObject];
                }) : element[objectIndex] = object[objectIndex];
        });
        return element;
    }

    clearMultipleItems(selected: any[]) {
        this.items = this.retrieveItemsFromCache();
        this.itemsCopy = this.retrieveCopyItemsFromCache();
        this.items.forEach((value, key) => {
            selected.forEach((selectedItem) => {
                if (JSON.stringify(this.changeObjectToArray(value)) === JSON.stringify(selectedItem)) {
                    this.items.delete(key);
                    this.itemsCopy.delete(key);
                }
            })
        });
        this.saveItemToCache(this.items);
        this.saveCopyItems(this.itemsCopy);
        window.location.reload();
    }

    clearItems() {
        this.items.clear();
        this.itemsCopy.clear();
        localStorage.setItem(this.itemName, JSON.stringify([]));
        localStorage.map = JSON.stringify(Array.from(this.items.entries()));
        localStorage.setItem('items-copy_' + this.itemName, JSON.stringify([]));
        localStorage.copy = JSON.stringify(Array.from(this.itemsCopy.entries()));
        window.location.reload();
    }

    removeItem(id) {
        this.items = this.retrieveItemsFromCache();
        this.itemsCopy = this.retrieveCopyItemsFromCache()
        this.items.forEach((value, key) => {
            if (JSON.stringify(this.changeObjectToArray(value)) === JSON.stringify(id)) {
                this.items.delete(key);
                this.itemsCopy.delete(key);
            }
        });
        this.saveItemToCache(this.items);
        this.saveCopyItems(this.itemsCopy);
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

    searchItems(index: number, element: Object, relativePath: string, itemName: string) {
        this._router.navigate([ relativePath + '/edit', index] );
        if (typeof element === 'string') {
            localStorage.setItem('edit', element);
            localStorage.setItem('edit-entity', itemName);
        }
    }

    editInheritedEntity(itemName: string, newElement: any) {
        this.itemName = itemName
        this.items = this.retrieveItemsFromCache();
        this.itemsCopy = this.retrieveCopyItemsFromCache();
        const editItem = JSON.parse(localStorage.getItem('edit'));

        if (JSON.stringify(newElement) !== JSON.stringify(editItem)) {
            this.items.forEach((currentMapValue, currentMapKey) => {
                Object.keys(currentMapValue).map((elementIndex) => {
                    if(currentMapValue[elementIndex] === JSON.stringify(editItem)) {
                        currentMapValue[elementIndex] = JSON.stringify(newElement);
                        const newId = this.itemName + JSON.stringify(currentMapValue);
                        this.items.delete(currentMapKey);
                        this.items.set(newId, currentMapValue);
                        this.itemsCopy.delete(currentMapKey);
                        this.itemsCopy.set(newId, currentMapValue);
                    }
                });
            });
        }
        this.saveItemToCache(this.items);
        this.saveCopyItems(this.itemsCopy);
    }

    editItems(index: number, element: Object) {
        this._router.navigate([this._router.url + '/edit' , index]);
        localStorage.setItem('edit', JSON.stringify(element));
    }

    checkRoute(element: Object) {
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

    loadSearchData() {
        const searchItem = JSON.parse(localStorage.getItem('search'));
        this.dataSources = new MatTableDataSource(this.changeMapToArray(this.retrieveItemsFromCache()));
        this.filterValue = searchItem;
        this.dataSources.filter = this.filterValue;
    }

    loadElement(element: Object) {
        const editItem = JSON.parse(localStorage.getItem('edit'));
        this.itemsCopy = this.retrieveCopyItemsFromCache();

        if (editItem !== null) {
            this.itemsCopy.forEach((currentValue) => {
                Object.keys(currentValue).map((elementIndex) => {
                    Object.keys(editItem).map((editIndex) => {
                        if(JSON.stringify(currentValue[editIndex]) === editItem[editIndex]
                            || currentValue[editIndex] === editItem[editIndex]) {
                            if (typeof currentValue[elementIndex] === 'string') {
                                currentValue[elementIndex].indexOf('-') !== -1 ?
                                    element[elementIndex] = new Date(currentValue[elementIndex])
                                    : element[elementIndex] = editItem[elementIndex];
                            } else {
                                typeof currentValue[elementIndex] === 'object' ?
                                    element[elementIndex] = currentValue[elementIndex]
                                    : element[elementIndex] = editItem[elementIndex];
                            }
                        }
                    })

                });
            })
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
