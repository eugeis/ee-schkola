import {Component, OnInit} from '@angular/core';

@Component({
    selector: 'app-template',
    templateUrl: './template.component.html',
    styleUrls: ['./template.component.scss']
})

export class TemplateComponent implements OnInit {

    street: string
    suite: string
    city: string
    code: string
    country: string
    index = 0;
    dataElement: any[][] = [];
    tempArray: any[][] = [];

    constructor() { }

    ngOnInit(): void {
    }

    inputElement() {
        this.street = this.simplifiedHtmlInputElement('street');
        this.suite = this.simplifiedHtmlInputElement('suite');
        this.city = this.simplifiedHtmlInputElement('city');
        this.code = this.simplifiedHtmlInputElement('code');
        this.country = this.simplifiedHtmlInputElement('country');
        this.dataElement.push([this.street, this.suite, this.city, this.code, this.country]);
    }

    deleteElement(index) {
        this.dataElement[index][0] = '';
        this.dataElement[index][1] = '';
        this.dataElement[index][2] = '';
        this.dataElement[index][3] = '';
        this.dataElement[index][4] = '';

        this.tempArray = this.dataElement.filter(function(element) {return element[0] !== '' })
        this.dataElement = this.tempArray;
    }

    printElement(index: number) {
        console.log('Street Value: ' + this.dataElement[index][0])
        console.log('Suite Value: ' + this.dataElement[index][1])
        console.log('City Value: ' + this.dataElement[index][2])
        console.log('Code Value: ' + this.dataElement[index][3])
        console.log('Country Value: ' + this.dataElement[index][4])
    }

    changeIndex(input: number) {
        this.index = input;
    }

    loadElement(index: number) {
        (<HTMLInputElement>document.getElementById('street')).value = this.dataElement[index][0];
        (<HTMLInputElement>document.getElementById('suite')).value = this.dataElement[index][1];
        (<HTMLInputElement>document.getElementById('city')).value = this.dataElement[index][2];
        (<HTMLInputElement>document.getElementById('code')).value = this.dataElement[index][3];
        (<HTMLInputElement>document.getElementById('country')).value = this.dataElement[index][4];
    }

    editElement(index: number) {
        this.dataElement[index][0] = this.simplifiedHtmlInputElement('street');
        this.dataElement[index][1] = this.simplifiedHtmlInputElement('suite');
        this.dataElement[index][2] = this.simplifiedHtmlInputElement('city');
        this.dataElement[index][3] = this.simplifiedHtmlInputElement('code');
        this.dataElement[index][4] = this.simplifiedHtmlInputElement('country');
    }

    simplifiedHtmlInputElement(element: string) {
        return (<HTMLInputElement>document.getElementById(element)).value;
    }
}
