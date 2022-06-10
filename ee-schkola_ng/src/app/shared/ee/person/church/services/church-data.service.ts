import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class ChurchDataService extends TableDataService {
    items = [];
    itemName = 'church';
}
