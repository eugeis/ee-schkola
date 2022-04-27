import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-location',
  templateUrl: './location.component.html',
  styleUrls: ['./location.component.scss']
})

export class LocationComponent implements OnInit {

  shelf: String
  fold: String
  index = 0;
  dataElement: any[][] = [];
  tempArray: any[][] = [];

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.shelf = this.simplifiedHtmlInputElement('shelf');
    this.fold = this.simplifiedHtmlInputElement('fold');
    this.dataElement.push([this.shelf, this.fold]);

  }
  deleteElement(index: number) {
    this.dataElement[index][0] = '';
    this.dataElement[index][1] = '';
    this.tempArray = this.dataElement.filter(function(element) {return element[0] !== '' })
    this.dataElement = this.tempArray;
  }
  printElement(index: number) {
    console.log('Shelf Value: ' + this.dataElement[index][0])
    console.log('Fold Value: ' + this.dataElement[index][1])
  }
  changeIndex(input: number) {
        this.index = input;
  }
  loadElement(index: number) {
    (<HTMLInputElement>document.getElementById('shelf')).value = this.dataElement[index][0];
    (<HTMLInputElement>document.getElementById('fold')).value = this.dataElement[index][1];
  }
  editElement(index: number) {
    this.dataElement[index][0] = this.simplifiedHtmlInputElement('shelf');
    this.dataElement[index][1] = this.simplifiedHtmlInputElement('fold');
  }
  simplifiedHtmlInputElement(element: string) {
        return (<HTMLInputElement>document.getElementById(element)).value;
  }

}



