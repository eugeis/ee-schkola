import {Book} from '@schkola/library/LibraryApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class BookDataService extends TableDataService {
    itemName = 'book';

    pageName = 'BookComponent';

    getFirst() {
        return new Book();
    }
}




