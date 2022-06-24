import {Injectable} from '@angular/core';
import {FormService} from '../../../template/services/form.service';

@Injectable()
export class ChurchViewService extends FormService {

    form = this.fb.group({
        name: '',
        address: '',
        association: '',
    });

    formArrayType = ['string', 'string', 'string'];

    formArrayName = ['Actions', 'name', 'address', 'association'];

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ChurchComponent';
}
