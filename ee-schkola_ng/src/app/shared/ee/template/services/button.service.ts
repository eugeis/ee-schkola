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

    // TODO: Change Map Key After Edit
    editElement(index: number, element: any) {
        const items = JSON.parse(localStorage.getItem(this.tableDataService.dataName));
        const elementWithNewDate = this.changeDateValue(element);
        items[index] = this.tableDataService.changeObjectToArray(elementWithNewDate);
        localStorage.setItem(this.tableDataService.dataName, JSON.stringify(items));
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
