import { ComponentFixture } from '@angular/core/testing';
import { TestBed, async } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from 'app/auth.service';

import { FeaturedComponent } from './featured.component';


  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, RouterTestingModule],
      declarations: [FeaturedComponent],
      providers: [AuthService],
    }).compileComponents();
  });

  

    describe('FeaturedComponent', () => {
  it('mounts', () => {
    // change the viewport of the test
    cy.viewport(1300, 700)
    // mount the component to test
    cy.mount(FeaturedComponent)
    // check visual text
   
    

    
  });
  
});
