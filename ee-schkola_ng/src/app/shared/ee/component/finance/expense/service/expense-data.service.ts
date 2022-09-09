import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class ExpenseViewComponent extends TableDataService {
    itemName = 'expense';

    pageName = 'ExpenseComponent';

    getFirst() {
        return new Expense();
    }
}




