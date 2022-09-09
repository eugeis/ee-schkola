import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class FeeKindViewComponent extends TableDataService {
    itemName = 'feekind';

    pageName = 'FeeKindComponent';

    getFirst() {
        return new FeeKind();
    }
}




