import {Book} from '@schkola/library/LibraryApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {BookDataService} from '../../service/book-data.service';

@Component({
  selector: 'app-book-list',
  templateUrl: './book-entity-list.component.html',
  styleUrls: ['./book-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: BookDataService}]
})

export class BookListComponent implements OnInit {

    book: Book = new Book();

    tableHeader: Array<String> = [];

    constructor(@Inject(BookDataService) public bookDataService: BookDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'title', 'description', 'language', 'releasedate', 'edition', 'category', 'author', 'location', 'id'];
    }
}




