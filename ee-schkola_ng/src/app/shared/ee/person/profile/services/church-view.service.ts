import {Injectable} from '@angular/core';
import {TemplateService} from '../../../template/services/template.service';

@Injectable()
export class ChurchViewService extends TemplateService {

    elementNameWithValue = [['name', 'string'], ['address', 'string'], ['association', 'string']];

    pageElement = ['Person'];

    pageName = 'ChurchComponent';
}

export const ELEMENT_DATA = [];

