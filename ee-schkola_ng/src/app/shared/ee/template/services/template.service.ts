import {Injectable} from '@angular/core';
import {FormBuilder, FormControl, Validators} from '@angular/forms';

@Injectable({ providedIn: 'root' })
export class TemplateService {

    protected elementNameWithValue: any[][];

    form = this.fb.group({});

    formArrayValue = this.fb.array([])

    formArrayName = [];

    formArrayType = [];

    formEnumElement = [];

    index = 0;

    constructor(private fb: FormBuilder) {
    }

    changeValue(input: string, elementName: string) {
        this.form.get(elementName).setValue(input);
    }

    init() {
        this.initElement(this.elementNameWithValue);
    }

    initElement(elements: any[][]) {
        elements.forEach(el => {
            const elementName = el.at(0);
            const elementType = el.at(1);
            if (elementType === 'string') {
                this.form.addControl(elementName, new FormControl('', [Validators.required]))
            } else if (elementType === 'number') {
                this.form.addControl(elementName, new FormControl(0, [Validators.required]))
            } else if (elementType === 'boolean') {
                this.form.addControl(elementName, new FormControl(false, [Validators.required]))
            } else if (elementType === 'datetime') {
                this.form.addControl(elementName, new FormControl('1.1.1970', [Validators.required]))
            } else if (elementType === 'enum') {
                const enumElement = el.at(2);
                let tempArray = [];
                Object.keys(enumElement).map((element, index) => {
                    tempArray.push(enumElement[index]);
                })
                tempArray = tempArray.filter((item) => item);
                this.formEnumElement.push(tempArray);
                this.form.addControl(elementName, new FormControl(0, [Validators.required]))
            } else {
                this.form.addControl(elementName, new FormControl('', [Validators.required]))
            }
            this.formArrayName.push(elementName);
            this.formArrayType.push(elementType);
        });
    }
}

export const ELEMENT_DATA = [];
