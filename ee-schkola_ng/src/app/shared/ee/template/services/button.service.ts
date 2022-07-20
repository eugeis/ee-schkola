import {Injectable} from '@angular/core';
import {TableDataService} from './data.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {

    constructor(private tableDataService: TableDataService) {
    }

    inputElement(element: any) {
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

    changeObjectStructure(object: Object) {
        const newObjectWithStructure: Object = {};
        Object.keys(object).map((objectIndex) => {
            typeof object[objectIndex] === 'object' ?
                object[objectIndex] instanceof Date ?
                    newObjectWithStructure[objectIndex] = this.formatDate(object[objectIndex]) :
                    Object.keys(object[objectIndex]).forEach((objectElementIndex) => {
                        newObjectWithStructure[objectElementIndex] = object[objectIndex][objectElementIndex];
                    }) : newObjectWithStructure[objectIndex] = object[objectIndex];
        })
        return newObjectWithStructure;
    }

    formatDate(date: Date) {
        return [date.getFullYear(), '0' + (date.getMonth() + 1), date.getDate()].join('-');
    }
}
