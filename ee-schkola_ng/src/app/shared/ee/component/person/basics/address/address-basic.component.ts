import {Address} from '../person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-address',
  templateUrl: './address.component.html',
  styleUrls: ['./address.component.scss'],
})

export class AddressComponent implements OnInit {

    @Input() address: Address = new Address();

    ngOnInit(): void {
        if (this.address === undefined) {
            this.address = {street: '', suite: '', city: '', code: '', country: ''};
        }
    }
}




