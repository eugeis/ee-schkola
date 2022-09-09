import {Book} from '../library/LibraryApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class BookViewComponent extends TableDataService {
    itemName = 'book';

    pageName = 'BookComponent';

    getFirst() {
        return new Book();
    }
}




