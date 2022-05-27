import {Injectable} from '@angular/core';
import {TemplateService} from '../../template/services/template.service';

@Injectable()
export class ProfileViewService extends TemplateService {

    elementNameWithValue = [['street', 'string'], ['suite', 'string'],
        ['city', 'string'], ['code', 'string'], ['country', 'string']];

}
