import {Book} from '../library/LibraryApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {BookDataService} from '../../services/book-data.service';

@Component({
  selector: 'app-book-list',
  templateUrl: './book-view.component.html',
  styleUrls: ['./book-view.component.scss'],
  providers: [{provide: TableDataService, useClass: BookDataService}]
})

export class BookViewComponent implements OnInit {

    book: Book = new Book();

    tableHeader: Array<String> = [];

    constructor(public bookDataService: BookDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'title', 'description', 'language', 'releasedate', 'edition', 'category', 'author', 'location', 'id'];
    }
}




