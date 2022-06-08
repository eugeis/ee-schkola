import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class ChurchDataService extends TableDataService {
    itemName = 'church';
}
