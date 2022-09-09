import {Book} from '../library/LibraryApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {BookDataService} from '../../services/book-data.service';

@Component({
  selector: 'app-book-view',
  templateUrl: './book-view.component.html',
  styleUrls: ['./book-view.component.scss'],
  providers: [{provide: TableDataService, useClass: BookDataService}]
})

export class BookViewComponent implements OnInit {


    book: Book;

    constructor(public bookDataService: BookDataService) {}

    ngOnInit(): void {
        this.book = this.bookDataService.getFirst();
        this.bookDataService.checkRoute(this.book);
    }
}




