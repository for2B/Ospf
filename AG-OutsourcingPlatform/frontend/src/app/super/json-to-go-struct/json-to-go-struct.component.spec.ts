import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { JsonToGoStructComponent } from './json-to-go-struct.component';

describe('JsonToGoStructComponent', () => {
  let component: JsonToGoStructComponent;
  let fixture: ComponentFixture<JsonToGoStructComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ JsonToGoStructComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(JsonToGoStructComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
