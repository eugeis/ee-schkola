import {Contact} from '@schkola/person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-contact',
  templateUrl: './contact-basic.component.html',
  styleUrls: ['./contact-basic.component.scss'],
})

export class ContactComponent implements OnInit {

    @Input() contact: Contact;
    
    ngOnInit() {
        if (this.contact === undefined) {
            this.contact = new Contact();
        }
        
    }
    
}




