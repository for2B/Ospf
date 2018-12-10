import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlatformProListComponent } from './platform-pro-list.component';

describe('PlatformProListComponent', () => {
  let component: PlatformProListComponent;
  let fixture: ComponentFixture<PlatformProListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlatformProListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlatformProListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
