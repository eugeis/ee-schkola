import {Injectable} from '@angular/core';
import {TemplateService} from '../../../template/services/template.service';

@Injectable()
export class PersonViewService extends TemplateService {

    pageElement = ['Person'];

    pageName = 'PersonComponent';
}

export const ELEMENT_DATA = [];

