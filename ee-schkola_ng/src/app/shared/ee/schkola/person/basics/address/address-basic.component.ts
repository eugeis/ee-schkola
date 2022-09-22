import {Address} from '@schkola/person/PersonApiBase'

import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-address',
  templateUrl: './address-basic.component.html',
  styleUrls: ['./address-basic.component.scss'],
})

export class AddressComponent {

    @Input() address: Address;
    
}




