import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-location',
  templateUrl: './location.component.html',
  styleUrls: ['./location.component.scss']
})

export class LocationComponent implements OnInit {

  shelf: String
  fold: String

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.shelf = (<HTMLInputElement>document.getElementById('shelf')).value;
    this.fold = (<HTMLInputElement>document.getElementById('fold')).value;
  }
  deleteElement() {
    this.shelf = '';
    this.fold = '';
  }
  printElement() {
    console.log('Shelf Value: ' + this.shelf)
    console.log('Fold Value: ' + this.fold)
  }

}



