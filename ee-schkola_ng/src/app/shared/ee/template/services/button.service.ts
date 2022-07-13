import {Injectable} from '@angular/core';
import {TableDataService} from './data.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {

    constructor(private tableDataService: TableDataService) {
    }

    inputElement(usedService: any) {
        usedService.formArrayValue.value.push(usedService.profile);
        this.tableDataService.addItemToTableArray(usedService.profile);
    }

    editElement(index: number, usedService: any) {
        for (const elementName in usedService.formArrayValue.value[index]) {
            if (usedService.formArrayValue.value[index].hasOwnProperty(elementName)) {
                usedService.formArrayValue.value[index][elementName] = usedService.form.get(elementName).value;
            }
        }
        localStorage.setItem(this.tableDataService.itemName, JSON.stringify(usedService.formArrayValue.value));
    }
}
