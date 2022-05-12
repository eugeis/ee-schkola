import {Template} from '../models/template';
import {Injectable} from '@angular/core';
import {TemplateService} from './template.service';

@Injectable({ providedIn: 'root' })
export class ButtonService {
    constructor(private templateService: TemplateService) {
    }

    // TODO: MAKE THIS DYNAMIC
    inputElement() {
        console.log(this.templateService.form.getRawValue());
        const element = this.templateService.form.getRawValue();
        this.templateService.formValue.push([element.street, element.suite,
            element.city, element.code, element.country, element.firstname, element.lastname]);
    }

    deleteElement(index) {
        for (const i in this.templateService.formValue[index]) {
            if (this.templateService.formValue[index].hasOwnProperty(i)) {
                this.templateService.formValue[index][i] = ''
            }
        }
        this.templateService.tempFormArray = this.templateService.formValue.filter(function(element) {return element[0] !== '' })
        this.templateService.formValue = this.templateService.tempFormArray;
    }

    printElement(index: number) {
        for (const i in this.templateService.formValue[index]) {
            if (this.templateService.formValue[index].hasOwnProperty(i)) {
                console.log(this.templateService.formArrayName[i] + ' Value: ' + this.templateService.formValue[index][i])
            }
        }
    }

    changeIndex(input: number) {
        this.templateService.index = input;
    }

    loadElement(index: number) {
        for (const i in this.templateService.formValue[index]) {
            if (this.templateService.formValue[index].hasOwnProperty(i)) {
                this.templateService.form.get(this.templateService.formArrayName[i]).setValue(this.templateService.formValue[index][i]);
            }
        }
    }

    editElement(index: number) {
        for (const i in this.templateService.formValue[index]) {
            if (this.templateService.formValue[index].hasOwnProperty(i)) {
                this.templateService.formValue[index][i] = this.templateService.form.get(this.templateService.formArrayName[i]).value;
            }
        }
    }
}
