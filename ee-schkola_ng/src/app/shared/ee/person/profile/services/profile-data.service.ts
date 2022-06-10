import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class ProfileDataService extends TableDataService {
    items = [];
    itemName = 'profile';
}
