import {Injectable} from '@angular/core';
import {TableDataService} from './data.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {

    constructor(private tableDataService: TableDataService) {
    }

    inputElement(element: any) {
        this.changeDateValue(element);
        const id = this.tableDataService.itemName + JSON.stringify(element);
        this.tableDataService.addItemToTableArray(element, id);
    }

    editElement(element: any) {
        this.tableDataService.items = this.tableDataService.retrieveItemsFromCache();
        this.tableDataService.tableItems = this.tableDataService.retrieveItemsForTableList();
        const editItem = JSON.parse(localStorage.getItem('edit'));
        const editItemEntity = localStorage.getItem('edit-entity');
        const newElement = this.changeDateValue(element);
        this.tableDataService.tableItems.forEach((currentValue, currentKey) => {
            const newId = this.tableDataService.itemName + JSON.stringify(newElement);
            if (this.removeSymbolFromString(JSON.stringify(editItem)).includes(
                this.removeSymbolFromString(JSON.stringify(this.tableDataService.changeObjectToArray(currentValue))))) {
                if (!this.removeSymbolFromString(JSON.stringify(this.tableDataService.items.get(currentKey))).includes(
                    this.removeSymbolFromString(JSON.stringify(newElement)))) {
                    this.tableDataService.items.delete(currentKey);
                    this.tableDataService.items.set(newId, newElement);
                    this.tableDataService.saveItemToCache(this.tableDataService.items);
                }
            }
        });
        this.tableDataService.editInheritedEntity(editItemEntity, newElement)
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
        return [date.getFullYear(), '0' + (date.getMonth() + 1), date.getDate()].join('-');
    }
}
