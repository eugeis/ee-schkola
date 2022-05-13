import {Component, OnInit} from '@angular/core';
import {TemplateService} from '../../services/template.service';
import {ChildTemplate} from '../../models/child-template';

@Component({
    selector: 'app-template',
    templateUrl: './template.component.html',
    styleUrls: ['./template.component.scss'],
    providers: [TemplateService],
})

export class TemplateComponent implements OnInit {

    elementNameWithValue = [['street', 'string'], ['suite', 'string'],
        ['city', 'string'], ['code', 'string'], ['country', 'string'], ['name', ChildTemplate]];
    childIndependent = false;
    constructor(public templateService: TemplateService) { }

    ngOnInit(): void {
        this.templateService.initElement(this.elementNameWithValue);
    }

    doSomething(test: any) {
        // this.templateService.initElementInterface(test);
    }
}
