import {Injectable} from '@angular/core';
import {TableDataService} from './data.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {

    constructor(private tableDataService: TableDataService) {
    }

    inputElement(element: any, entityElements: Array<string>) {
        this.changeDateValue(element);
        const id = this.tableDataService.itemName + JSON.stringify(element);
        this.tableDataService.addItemToTableArray(element, id, entityElements);
    }

    editElement(element: any) {
        const mapItems = this.tableDataService.retrieveItemsFromCache();
        const itemsCopy = this.tableDataService.retrieveCopyItemsFromCache();
        const editItem = JSON.parse(localStorage.getItem('edit'));
        const editItemEntity = localStorage.getItem('edit-entity');
        const newElement = this.changeDateValue(element);
        if (JSON.stringify(newElement) !== JSON.stringify(editItem)) {
            mapItems.forEach((currentValue, currentKey) => {
                const newId = this.tableDataService.itemName + JSON.stringify(newElement);
                if (JSON.stringify(this.tableDataService.changeObjectToArray(currentValue)) === JSON.stringify(editItem)) {
                    mapItems.set(newId, newElement);
                    mapItems.delete(currentKey);
                    itemsCopy.delete(currentKey);
                    itemsCopy.set(newId, currentValue);
                }
            });
        }
        this.tableDataService.saveItemToCache(mapItems);
        this.tableDataService.saveCopyItems(itemsCopy);
        this.tableDataService.editInheritedEntity(editItemEntity, element)
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
