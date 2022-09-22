import {Fee} from '@schkola/finance/FinanceApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class FeeDataService extends TableDataService {
    itemName = 'fee';

    pageName = 'FeeComponent';

    getFirst() {
        return new Fee();
    }
}




