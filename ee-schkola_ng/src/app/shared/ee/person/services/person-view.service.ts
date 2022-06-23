import {Injectable} from '@angular/core';
import {FormService} from '../../template/services/form.service';

@Injectable()
export class PersonViewService extends FormService {

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'PersonComponent';
}
