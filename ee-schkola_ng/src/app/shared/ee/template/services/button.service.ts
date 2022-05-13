import {Injectable} from '@angular/core';
import {TemplateService} from './template.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {
    constructor(private templateService: TemplateService) {
    }

    inputElement() {
        this.templateService.formArrayValue.value.push(this.templateService.form.getRawValue());
    }

    deleteElement(index) {
        this.templateService.formArrayValue.removeAt(index)
    }

    printElement(index: number) {
        for (const elementName in this.templateService.formArrayValue.value[index]) {
            if (this.templateService.formArrayValue.value[index].hasOwnProperty(elementName)) {
                console.log(elementName + ' Value: ' + this.templateService.formArrayValue.value[index][elementName])
            }
        }
    }

    changeIndex(input: number) {
        this.templateService.index = input;
    }

    loadElement(index: number) {
        for (const elementName in this.templateService.formArrayValue.value[index]) {
            if (this.templateService.formArrayValue.value[index].hasOwnProperty(elementName)) {
                this.templateService.form.get(elementName).setValue(this.templateService.formArrayValue.value[index][elementName]);
            }
        }
    }

    editElement(index: number) {
        for (const elementName in this.templateService.formArrayValue.value[index]) {
            if (this.templateService.formArrayValue.value[index].hasOwnProperty(elementName)) {
                this.templateService.formArrayValue.value[index][elementName] = this.templateService.form.get(elementName).value;
            }
        }
    }
}
