import {Component, OnInit} from '@angular/core';
import {FormBuilder, Validators} from '@angular/forms';
import {Template} from '../../models/template';

@Component({
    selector: 'app-template',
    templateUrl: './template.component.html',
    styleUrls: ['./template.component.scss'],
})

export class TemplateComponent implements OnInit {

    form = this.fb.group({
        street: ['', [Validators.required]],
        suite: ['', [Validators.required]],
        city: ['', [Validators.required]],
        code: ['', [Validators.required]],
        country: ['', [Validators.required]]
    });
    formArrayName = ['street', 'suite', 'city', 'code', 'country'];
    index = 0;
    formValue: any[][] = [];
    tempFormArray: any[][] = [];

    constructor(private fb: FormBuilder) { }

    ngOnInit(): void {
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
