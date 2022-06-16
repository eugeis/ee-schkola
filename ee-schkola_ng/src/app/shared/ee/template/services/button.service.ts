import {Injectable} from '@angular/core';
import {TemplateService} from './template.service';
import {TableDataService} from './data.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {

    constructor(private templateService: TemplateService, private tableDataService: TableDataService) {
    }

    inputElement() {
        this.templateService.formArrayValue.value.push(this.templateService.form.getRawValue());
        this.tableDataService.addItemToTableArray(this.templateService.form.getRawValue());
        console.log(this.templateService.form.getRawValue());
    }

    editElement(index: number) {
        this.tableDataService.retrieveItemFromCache();
        this.tableDataService.items.forEach((el) => {
            this.templateService.formArrayValue.value.push(el);
        })
        for (const elementName in this.templateService.formArrayValue.value[index]) {
            if (this.templateService.formArrayValue.value[index].hasOwnProperty(elementName)) {
                this.templateService.formArrayValue.value[index][elementName] = this.templateService.form.get(elementName).value;
            }
        }
        localStorage.setItem(this.tableDataService.itemName, JSON.stringify(this.templateService.formArrayValue.value));
    }
}
