import {Template} from '../models/template';
import {Injectable} from '@angular/core';
import {FormBuilder, FormControl, Validators} from '@angular/forms';

@Injectable({ providedIn: 'root' })
export class TemplateService {

    form = this.fb.group({});

    formArrayName = [];

    index = 0;

    formValue: any[][] = [];

    tempFormArray: any[][] = [];

    constructor(private fb: FormBuilder) {
    }

    initElement(elements: any[][]) {
        elements.forEach(el => {
            const elementType = el.at(1);
            const elementName = el.at(0);
            if (elementType === 'string') {
                this.form.addControl(elementName, new FormControl('', [Validators.required]))
            } else if (elementType === 'number') {
                this.form.addControl(elementName, new FormControl(0, [Validators.required]))
            } else if (elementType === 'boolean') {
                this.form.addControl(elementName, new FormControl(false, [Validators.required]))
            } else {
                this.form.addControl(elementName, new FormControl('', [Validators.required]))
            }
            this.formArrayName.push(elementName);
        })
    }

    inputElement() {
        const element = this.form.getRawValue() as Template;
        this.formValue.push([element.street, element.suite, element.city, element.code, element.country]);
    }

    deleteElement(index) {
        for (const i in this.formValue[index]) {
            if (this.formValue[index].hasOwnProperty(i)) {
                this.formValue[index][i] = ''
            }
        }
        this.tempFormArray = this.formValue.filter(function(element) {return element[0] !== '' })
        this.formValue = this.tempFormArray;
    }

    printElement(index: number) {
        for (const i in this.formValue[index]) {
            if (this.formValue[index].hasOwnProperty(i)) {
                console.log(this.formArrayName[i] + ' Value: ' + this.formValue[index][i])
            }
        }
    }

    changeIndex(input: number) {
        this.index = input;
    }

    loadElement(index: number) {
        for (const i in this.formValue[index]) {
            if (this.formValue[index].hasOwnProperty(i)) {
                this.form.get(this.formArrayName[i]).setValue(this.formValue[index][i]);
            }
        }
    }

    editElement(index: number) {
        for (const i in this.formValue[index]) {
            if (this.formValue[index].hasOwnProperty(i)) {
                this.formValue[index][i] = this.form.get(this.formArrayName[i]).value;
            }
        }
    }
}
