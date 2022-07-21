import {Injectable} from '@angular/core';
import {TableDataService} from './data.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {

    constructor(private tableDataService: TableDataService) {
    }

    inputElement(element: any) {
        element['birthday'] = this.formatDate(element['birthday']);
        const id = element['birthName'] + element['gender'];
        this.tableDataService.addItemToTableArray(element, id);
    }

    // TODO: Implement Edit Element with Current Approach
    editElement(index: number, usedService: any) {
        for (const elementName in usedService.formArrayValue.value[index]) {
            if (usedService.formArrayValue.value[index].hasOwnProperty(elementName)) {
                usedService.formArrayValue.value[index][elementName] = usedService.form.get(elementName).value;
            }
        }
        localStorage.setItem(this.tableDataService.itemName, JSON.stringify(usedService.formArrayValue.value));
    }

    formatDate(date: Date) {
        return [date.getFullYear(), '0' + (date.getMonth() + 1), date.getDate()].join('-');
    }
}
