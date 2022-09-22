import {Book, Location} from '@schkola/library/LibraryApiBase'
import {PersonName} from '@schkola/person/PersonApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {BookDataService} from '../../service/book-data.service';

@Component({
  selector: 'app-book-view',
  templateUrl: './book-entity-view.component.html',
  styleUrls: ['./book-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: BookDataService}]
})

export class BookViewComponent implements OnInit {


    book: Book;

    constructor(@Inject(BookDataService) public bookDataService: BookDataService) {}

    ngOnInit(): void {
        this.book = this.bookDataService.getFirst();
        this.book.author = new PersonName();
        this.book.location = new Location();
        this.bookDataService.checkRoute(this.book);
    }
}




