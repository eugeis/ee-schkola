import {Book, Location} from '@schkola/library/LibraryApiBase'
import {PersonName} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {BookDataService} from '../../service/book-data.service';

@Component({
  selector: 'app-book-form',
  templateUrl: './book-form.component.html',
  styleUrls: ['./book-form.component.scss']
})

export class BookFormComponent implements OnInit {


    @Input() book: Book;

    constructor(public bookDataService: BookDataService) {}

    ngOnInit(): void {
        if (this.book.author === undefined) {
            this.book.author = new PersonName();
        }
        if (this.book.location === undefined) {
            this.book.location = new Location();
        }
    }
}




