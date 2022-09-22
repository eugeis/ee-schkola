import {Contact} from '@schkola/person/PersonApiBase'

import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-contact',
  templateUrl: './contact-basic.component.html',
  styleUrls: ['./contact-basic.component.scss'],
})

export class ContactComponent {

    @Input() contact: Contact;
    
}




