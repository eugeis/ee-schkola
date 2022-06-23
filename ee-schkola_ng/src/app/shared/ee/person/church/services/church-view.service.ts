import {Injectable} from '@angular/core';
import {FormService} from '../../../template/services/form.service';

@Injectable()
export class ChurchViewService extends FormService {

    elementNameWithValue = [['name', 'string'], ['address', 'string'], ['association', 'string']];

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ChurchComponent';

    formArrayName = ['Actions'];
}
