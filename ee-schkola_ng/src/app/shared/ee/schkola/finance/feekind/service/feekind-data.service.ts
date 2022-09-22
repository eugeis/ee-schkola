import {FeeKind} from '@schkola/finance/FinanceApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class FeeKindDataService extends TableDataService {
    itemName = 'feekind';

    pageName = 'FeeKindComponent';

    getFirst() {
        return new FeeKind();
    }
}




