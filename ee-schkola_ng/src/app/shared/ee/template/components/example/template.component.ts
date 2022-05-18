import {Component, Input, OnInit} from '@angular/core';
import {TemplateService} from '../../services/template.service';

@Component({
    selector: 'app-template',
    templateUrl: './template.component.html',
    styleUrls: ['./template.component.scss'],
    providers: [TemplateService],
})

export class TemplateComponent implements OnInit {

    @Input() showElement: boolean;

    elementNameWithValue = [['street', 'string'], ['suite', 'string'],
        ['city', 'string'], ['code', 'string'], ['country', 'string']];
    show = false;

    constructor(public templateService: TemplateService) { }

    ngOnInit(): void {
        this.show = this.showElement;
        this.templateService.initElement(this.elementNameWithValue);
    }
}
