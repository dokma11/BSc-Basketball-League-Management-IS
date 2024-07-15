import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProposeTradeFormComponent } from './propose-trade-form.component';

describe('ProposeTradeFormComponent', () => {
  let component: ProposeTradeFormComponent;
  let fixture: ComponentFixture<ProposeTradeFormComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ProposeTradeFormComponent]
    });
    fixture = TestBed.createComponent(ProposeTradeFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
