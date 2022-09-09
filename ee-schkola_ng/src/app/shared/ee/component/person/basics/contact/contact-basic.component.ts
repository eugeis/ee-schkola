import {Contact} from '../person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.scss'],
})

export class ContactComponent implements OnInit {

    @Input() contact: Contact = new Contact();

    ngOnInit(): void {
        if (this.contact === undefined) {
            this.contact = {phone: '', email: '', cellphone: ''};
        }
    }
}




