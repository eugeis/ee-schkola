import {Injectable} from '@angular/core';
import {FormBuilder, FormControl, Validators} from '@angular/forms';

@Injectable({ providedIn: 'root' })
export class TemplateService {

    form = this.fb.group({});

    formArrayValue = this.fb.array([])

    formArrayName = [];

    formArrayChild = [];

    index = 0;

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
                console.log('HERE!');
                console.log(this.formArrayChild.length);
                this.formArrayChild.forEach(element => {
                    console.log('ELEMENT: ' + element);
                })
                this.form.addControl(elementName, new FormControl('', [Validators.required]))
            }
            this.formArrayName.push(elementName);
        });
    }

    initElementInterface(elements: any[][]) {
        this.formArrayChild = elements;
    }
}
