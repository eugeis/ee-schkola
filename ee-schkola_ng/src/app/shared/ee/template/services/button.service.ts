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

    editElement(index: number, element: any) {
        const mapItems = this.tableDataService.retrieveItemFromCache();
        const arrayItems = JSON.parse(localStorage.getItem(this.tableDataService.dataName));
        const newElement = this.changeDateValue(element);
        mapItems.forEach((currentValue, currentKey) => {
            const newId = this.tableDataService.itemName + JSON.stringify(newElement);
            if (JSON.stringify(this.tableDataService.changeObjectToArray(currentValue)) === JSON.stringify(arrayItems[index])) {
                mapItems.set(newId, newElement);
                mapItems.delete(currentKey);
            }
        });
        arrayItems[index] = this.tableDataService.changeObjectToArray(newElement);
        localStorage.setItem(this.tableDataService.dataName, JSON.stringify(arrayItems));
        localStorage.map = JSON.stringify(Array.from(mapItems.entries()));
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
