import {Address} from '@schkola/person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-address',
  templateUrl: './address-basic.component.html',
  styleUrls: ['./address-basic.component.scss'],
})

export class AddressComponent implements OnInit {

    @Input() address: Address;
    
    ngOnInit() {
        if (this.address === undefined) {
            this.address = new Address();
        }
        
    }
    
}




